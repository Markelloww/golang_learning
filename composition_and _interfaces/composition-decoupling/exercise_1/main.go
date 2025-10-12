package main

import "fmt"

type administrator interface {
	administrate(system string)
}

type developer interface {
	develop(system string)
}

type adminlist struct {
	list []administrator
}

func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

type devlist struct {
	list []developer
}

func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

type sysadmin struct {
	name string
}

func (sa *sysadmin) administrate(system string) {
	fmt.Println(sa.name, "is admining", system)
}

type programmer struct {
	name string
}

func (proger *programmer) develop(system string) {
	fmt.Println(proger.name, "is developing", system)
}

type company struct {
	administrator
	developer
}

func main() {
	var admins adminlist
	var devs devlist

	admins.Enqueue(&sysadmin{"Oleg"})

	devs.Enqueue(&programmer{"Mark"})
	devs.Enqueue(&programmer{"Nikira"})

	comp := company{
		administrator: admins.Dequeue(),
		developer:     devs.Dequeue(),
	}

	admins.Enqueue(&comp)
	devs.Enqueue(&comp)

	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	for _, task := range tasks {
		if task.needsAdmin {
			adm := admins.Dequeue()
			adm.administrate(task.system)
			continue
		}
		dev := devs.Dequeue()
		dev.develop(task.system)
	}
}
