package gen

type (
	// Item is the interface for anything which can produce a file.
	Item interface {

		// String gets the file name or package name without any extension,
		// e.g. `blockchain`, not `blockchain.dart`
		String() string

		// Write will write the file for this item in the given base path.
		// This base path is the folder that should contain the `lib` folder.
		Write(dryRun bool, basePath, packageName string)
	}

	// Node is the interface for any part of the dependency tree being built.
	Node interface {
		Item

		// Children will return any dependencies for this node.
		Children() []Node

		// Group is the collection of items which share a folder or library.
		Group() *Group

		// setGroup should only be called by a group when an item is being added to the group
		// and is used to set the group an item belongs to.
		setGroup(group *Group, index int)

		// PrintTree prints this node and any children to the console as a tree.
		PrintTree(indent string, last bool)
	}
)
