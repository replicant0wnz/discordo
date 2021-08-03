package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rigormorrtiss/discordo/util"
	"github.com/rivo/tview"
)

func NewChannelsTreeView(channelsTreeNode *tview.TreeNode, onChannelsTreeViewSelected func(node *tview.TreeNode), theme *util.Theme) *tview.TreeView {
	channelsTreeView := tview.NewTreeView()

	channelsTreeView.
		SetTopLevel(1).
		SetRoot(channelsTreeNode).
		SetCurrentNode(channelsTreeNode).
		SetSelectedFunc(onChannelsTreeViewSelected).
		SetBackgroundColor(tcell.GetColor(theme.TreeViewBackground)).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	return channelsTreeView
}