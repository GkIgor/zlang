package target

var target string = "./default.etq"

func GetTarget() string {
	return target
}

func SetTarget(t string) {
	target = t
}
