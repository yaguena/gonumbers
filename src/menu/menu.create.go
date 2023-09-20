package menu

import "github.com/dixonwille/wmenu"

func execute(opt wmenu.Opt) error {
	cmd, err := opt.Value.(NameEntity)
	if !err {
		log.Fatal("Could not cast option's value to NameEntity")
	}
	fmt.Printf("%s has an id of %d.\n", opt.Text, opt.ID)
	fmt.Printf("Hello, %s %s.\n", name.FirstName, name.Lastwmenu
	return nil
}

func Create() {
	menu := NewMenu("Choose an option.")
	menu.ChangeReaderWriter(reader, os.Stdout, os.Stderr)
	menu.Action(execute)
	menu.Option("Option 0", NameEntity{"Bill", "Bob"}, true, nil)
	menu.Option("Option 1", NameEntity{"John", "Doe"}, false, nil)
	menu.Option("Option 2", NameEntity{"Jane", "Doe"}, false, nil)
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
