package tree

import "fmt"

type dependecyTreeNode struct {
	name string
	dependencies []*dependecyTreeNode
}

type dependencyTree struct {
	projects []*dependecyTreeNode
	projectToNodeMap map[string]*dependecyTreeNode
}

func InitDependencyManager() *dependencyTree{
	d := dependencyTree{}
	d.projects = make([]*dependecyTreeNode, 0)
	d.projectToNodeMap = make(map[string]*dependecyTreeNode)
	return &d
}

func (d *dependencyTree) AddDependency(project, dependentOn string) {
	projectNode := d.AddProjectGetRef(project)
	dependency := d.AddProjectGetRef(dependentOn)
	projectNode.dependencies = append(projectNode.dependencies, dependency)
}

func (d *dependencyTree) AddProjectGetRef(project string) *dependecyTreeNode{
	var projectNode *dependecyTreeNode
	if _, ok := d.projectToNodeMap[project]; ok {
		projectNode = d.projectToNodeMap[project]
	}else {
		projectNode = &dependecyTreeNode{
			name:         project,
			dependencies: make([]*dependecyTreeNode, 0),
		}
		d.projectToNodeMap[project] = projectNode
		d.projects = append(d.projects, projectNode)
	}
	return projectNode
}

func (d dependencyTree) GetBuildOrder() []string{
	visited := make(map[string]bool)
	for _, p := range d.projects{
		visited[p.name] = false
	}
	result := make([]string, 0)

	for _, project := range d.projects{
		postOrderProcessing(project, visited, &result)
	}
	return result

}

func postOrderProcessing(root *dependecyTreeNode, visited map[string]bool, result *[]string){
	if visited[root.name]{
		return
	}
	for _, val := range root.dependencies{
		postOrderProcessing(val, visited, result)
	}
	visited[root.name] = true
	*result = append(*result, root.name)
}

func PlayDependencyTree(){
	d := InitDependencyManager()
	d.AddDependency("a", "d")
	d.AddDependency("f", "b")
	d.AddDependency("b", "d")
	d.AddDependency("f", "a")
	d.AddDependency("d", "c")
	r := d.GetBuildOrder()
	fmt.Println(r)
}