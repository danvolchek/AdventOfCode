package vm

func add(vm *VM, arg0, arg1, arg2 arg) {
	arg2.write(arg0.read() + arg1.read())
}

func mul(vm *VM, arg0, arg1, arg2 arg) {
	arg2.write(arg0.read() * arg1.read())
}

func in(vm *VM, arg0 arg) {
	arg0.write(vm.read())
}

func out(vm *VM, arg0 arg) {
	vm.write(arg0.read())
}

func jit(vm *VM, arg0, arg1 arg) {
	if arg0.read() != 0 {
		vm.PC = arg1.read()
	}
}

func jif(vm *VM, arg0, arg1 arg) {
	if arg0.read() == 0 {
		vm.PC = arg1.read()
	}
}

func lt(vm *VM, arg0, arg1, arg2 arg) {
	val := 0
	if arg0.read() < arg1.read() {
		val = 1
	}
	arg2.write(val)
}

func eq(vm *VM, arg0, arg1, arg2 arg) {
	val := 0
	if arg0.read() == arg1.read() {
		val = 1
	}
	arg2.write(val)
}
