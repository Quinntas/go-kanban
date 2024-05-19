package main

import "github.com/charmbracelet/bubbles/list"

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}
	b.cols[todo].list.Title = "Todo"
	b.cols[todo].list.SetItems([]list.Item{
		NewTask(todo, "write kanban board", "it's a demo, so don't worry"),
		NewTask(todo, "write documentation", "it's a demo, so don't worry"),
		NewTask(todo, "write tests", "it's a demo, so don't worry"),
	})
	b.cols[inProgress].list.Title = "In Progress"
	b.cols[inProgress].list.SetItems([]list.Item{
		NewTask(inProgress, "write kanban board", "it's a demo, so don't worry"),
		NewTask(inProgress, "write documentation", "it's a demo, so don't worry"),
	})
	b.cols[done].list.Title = "Done"
	b.cols[done].list.SetItems([]list.Item{
		NewTask(done, "write kanban board", "it's a demo, so don't worry"),
	})
}
