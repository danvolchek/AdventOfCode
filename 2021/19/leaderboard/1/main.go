package main

// Not solved yet

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "19", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type pos struct {
	x, y, z int
}

func (p pos) String() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y) + "," + strconv.Itoa(p.z)
}

func (p pos) sub(o pos) pos {
	return pos{
		x: p.x - o.x,
		y: p.y - o.y,
		z: p.z - o.z,
	}
}

func (p pos) add(o pos) pos {
	return pos{
		x: p.x + o.x,
		y: p.y + o.y,
		z: p.z + o.z,
	}
}

func parse(r io.Reader) [][]pos {
	raw, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rawScanners := bytes.Split(raw, []byte{'\n', '\n'})

	var scanners [][]pos
	for _, rawScanner := range rawScanners {

		var beacons []pos

		rawBeacons := bytes.Split(bytes.TrimSpace(rawScanner), []byte{'\n'})
		for _, rawBeacon := range rawBeacons {
			parts := bytes.Split(rawBeacon, []byte{','})
			if len(parts) != 3 {
				continue
			}

			beacons = append(beacons, pos{
				x: toInt(parts[0]),
				y: toInt(parts[1]),
				z: toInt(parts[2]),
			})

		}

		scanners = append(scanners, beacons)
	}

	return scanners
}

func toInt(num []byte) int {
	v, err := strconv.Atoi(string(num))
	if err != nil {
		panic(err)
	}

	return v
}

func solve(r io.Reader) {
	scanners := parse(r)

	scannerPositions := map[int]pos{
		0: {x: 0, y: 0, z: 0},
	}

	tried := make(map[try]bool)
	for len(scannerPositions) != len(scanners) {
		changed := normalize(scanners, scannerPositions, tried)
		if !changed {
			panic("failed to normalize beacon list")
		}
	}

	fmt.Printf("%+v\n", count(scanners, scannerPositions))
}

func count(scanners [][]pos, scannerPositions map[int]pos) int {
	beaconSet := make(map[pos]bool)
	for i, beacons := range scanners {
		for _, beacon := range beacons {
			beaconSet[beacon.add(scannerPositions[i])] = true
		}
	}

	/*for becaon := range beaconSet {
		fmt.Println(becaon)
	}*/

	return len(beaconSet)
}

type try struct {solved, unsolved int}

func normalize(scanners [][]pos, scannerPositions map[int]pos, tried map[try]bool) bool {
	for solvedIndex, solvedBeacons := range scanners {
		solvedPosition, ok := scannerPositions[solvedIndex]
		if !ok {
			continue
		}

		for unsolvedIndex, unsolvedBeacons := range scanners {
			if _, ok := scannerPositions[unsolvedIndex]; ok {
				continue
			}

			thisTry := try{solved: solvedIndex, unsolved:unsolvedIndex}
			if tried[thisTry] {
				continue
			}

			fmt.Printf("Trying to solve %d using %d\n", unsolvedIndex, solvedIndex)

			for _, variation := range magnetize(unsolvedBeacons) {
				if tRef, jRef, ok := beaconsOverlap(solvedBeacons, variation); ok {
					scanners[unsolvedIndex] = variation

					unsolvedPosition := solvedPosition.add(tRef.sub(jRef))
					fmt.Printf("Solved %d using %d, doing %v + %v - %v = %v\n", unsolvedIndex, solvedIndex, solvedPosition, tRef, jRef, unsolvedPosition)
					scannerPositions[unsolvedIndex] = unsolvedPosition // something to do with solvedPosition, tRef, jRef

					return true
				}
			}

			tried[thisTry] = true
		}
	}

	return false
}

func magnetize(scanner []pos) [48][]pos {
	var ret [48][]pos
	i := 0

	for _, xMul := range []int{1, -1} {
		for _, yMul := range []int{1, -1} {
			for _, zMul := range []int{1, -1} {
				for _, shiftAmnt := range []int{0, 1, 2} {
					for _, swapAmnt := range []bool{true, false} {
						ret[i] = buildBeacons(scanner, chain(swap(swapAmnt), shift(shiftAmnt), mul(xMul, yMul, zMul)))
						i++
					}
				}
			}
		}
	}

	return ret
}

func chain(funcs ...func(pos) pos) func(pos) pos {
	return func(p pos) pos {
		for _, f := range funcs {
			p = f(p)
		}

		return p
	}
}

func mul(x, y, z int) func(pos) pos {
	return func(p pos) pos {
		return pos{
			x: p.x * x,
			y: p.y * y,
			z: p.z * z,
		}
	}
}

func swap(doSwap bool) func(pos)pos {
	return func(p pos) pos {
		if doSwap {
			return pos{
				x: p.y,
				y: p.x,
				z: p.z,
			}
		}
		return p
	}
}

// 1 2 3
// 3 1 2 right 1
// 2 3 1 right 2


// 2 1 3
// 3 2 1
// 1 3 2

func shift(amnt int) func(pos) pos {
	return func(p pos) pos {
		for i := 0; i < amnt; i++ {
			p = pos{
				x: p.y,
				y: p.z,
				z: p.x,
			}
		}

		return p
	}
}

func buildBeacons(input []pos, transmute func(pos) pos) []pos {
	ret := make([]pos, len(input))
	for i, j := range input {
		ret[i] = transmute(j)
	}

	return ret
}

func beaconsOverlap(target, check []pos) (pos, pos, bool) {
	for _, tRef := range target {
		for _, jref := range check {
			align1, align2 := alignment(tRef, target), alignment(jref, check)

			if alignmentsMatch(align1, align2) {
				return tRef, jref, true
			}
		}
	}

	return pos{}, pos{}, false
}

func alignment(target pos, values []pos) []pos {
	ret := make([]pos, len(values))
	for i, v := range values {
		ret[i] = v.sub(target)
	}

	return ret
}

func alignmentsMatch(target, check []pos) bool {
	matching := 0
	for _, beacon := range target {
		if contains(beacon, check) {
			matching += 1
		}
	}

	return matching >= 12
}

func contains(i pos, j []pos) bool {
	for _, jj := range j {
		if i == jj {
			return true
		}
	}

	return false
}

func main() {
	//solve(strings.NewReader("--- scanner 0 ---\n0,2\n4,1\n3,3\n\n--- scanner 1 ---\n-1,-1\n-5,0\n-2,1"))
	//solve(strings.NewReader("--- scanner 0 ---\n404,-588,-901\n528,-643,409\n-838,591,734\n390,-675,-793\n-537,-823,-458\n-485,-357,347\n-345,-311,381\n-661,-816,-575\n-876,649,763\n-618,-824,-621\n553,345,-567\n474,580,667\n-447,-329,318\n-584,868,-557\n544,-627,-890\n564,392,-477\n455,729,728\n-892,524,684\n-689,845,-530\n423,-701,434\n7,-33,-71\n630,319,-379\n443,580,662\n-789,900,-551\n459,-707,401\n\n--- scanner 1 ---\n686,422,578\n605,423,415\n515,917,-361\n-336,658,858\n95,138,22\n-476,619,847\n-340,-569,-846\n567,-361,727\n-460,603,-452\n669,-402,600\n729,430,532\n-500,-761,534\n-322,571,750\n-466,-666,-811\n-429,-592,574\n-355,545,-477\n703,-491,-529\n-328,-685,520\n413,935,-424\n-391,539,-444\n586,-435,557\n-364,-763,-893\n807,-499,-711\n755,-354,-619\n553,889,-390\n\n--- scanner 2 ---\n649,640,665\n682,-795,504\n-784,533,-524\n-644,584,-595\n-588,-843,648\n-30,6,44\n-674,560,763\n500,723,-460\n609,671,-379\n-555,-800,653\n-675,-892,-343\n697,-426,-610\n578,704,681\n493,664,-388\n-671,-858,530\n-667,343,800\n571,-461,-707\n-138,-166,112\n-889,563,-600\n646,-828,498\n640,759,510\n-630,509,768\n-681,-892,-333\n673,-379,-804\n-742,-814,-386\n577,-820,562\n\n--- scanner 3 ---\n-589,542,597\n605,-692,669\n-500,565,-823\n-660,373,557\n-458,-679,-417\n-488,449,543\n-626,468,-788\n338,-750,-386\n528,-832,-391\n562,-778,733\n-938,-730,414\n543,643,-506\n-524,371,-870\n407,773,750\n-104,29,83\n378,-903,-323\n-778,-728,485\n426,699,580\n-438,-605,-362\n-469,-447,-387\n509,732,623\n647,635,-688\n-868,-804,481\n614,-800,639\n595,780,-596\n\n--- scanner 4 ---\n727,592,562\n-293,-554,779\n441,611,-461\n-714,465,-776\n-743,427,-804\n-660,-479,-426\n832,-632,460\n927,-485,-438\n408,393,-506\n466,436,-512\n110,16,151\n-258,-428,682\n-393,719,612\n-211,-452,876\n808,-476,-593\n-575,615,604\n-485,667,467\n-680,325,-822\n-627,-443,-432\n872,-547,-609\n833,512,582\n807,604,487\n839,-516,451\n891,-625,532\n-652,-548,-490\n30,-46,-14"))
	solve(input())
}
