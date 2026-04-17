package main

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	res := args[0]
	for i := 1; i < len(args); i++{
		res += "\n" + args[i]
	}
	return res
}