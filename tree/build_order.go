package tree

import "fmt"

type dependecyTreeNode struct {
	name         string
	dependencies []*dependecyTreeNode
}

type dependencyTree struct {
	projects         []*dependecyTreeNode
	projectToNodeMap map[string]*dependecyTreeNode
}

func InitDependencyManager() *dependencyTree {
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

func (d *dependencyTree) AddProjectGetRef(project string) *dependecyTreeNode {
	var projectNode *dependecyTreeNode
	if _, ok := d.projectToNodeMap[project]; ok {
		projectNode = d.projectToNodeMap[project]
	} else {
		projectNode = &dependecyTreeNode{
			name:         project,
			dependencies: make([]*dependecyTreeNode, 0),
		}
		d.projectToNodeMap[project] = projectNode
		d.projects = append(d.projects, projectNode)
	}
	return projectNode
}

func (d dependencyTree) GetBuildOrder() []string {
	visited := make(map[string]bool)
	for _, p := range d.projects {
		visited[p.name] = false
	}
	result := make([]string, 0)

	for _, project := range d.projects {
		postOrderProcessing(project, visited, &result)
	}
	return result

}

func (d *dependencyTree) checkCycle() (bool, []string) {
	unvisited := make(map[string]bool)
	currentStack := make(map[string]bool)
	completed := make(map[string]bool)
	for _, val := range d.projects {
		unvisited[val.name] = true
	}
	for _, val := range d.projects {
		var cycle []string
		result, cycle := d.dfsForCycle(val.name, unvisited, currentStack, completed, &cycle)
		if result {
			return result, cycle
		}
	}

	return false, []string{}
}

func (d *dependencyTree) dfsForCycle(project string, unvisited, currentStack, completed map[string]bool, cycle *[]string) (bool, []string) {
	if _, ok := completed[project]; ok {
		return false, *cycle
	}
	*cycle = append(*cycle, project)
	vertex := d.projectToNodeMap[project]
	if _, ok := unvisited[project]; ok {
		delete(unvisited, project)
	}
	if _, ok := currentStack[project]; ok {
		return true, *cycle
	} else {
		currentStack[project] = true
	}
	for _, dependency := range vertex.dependencies {
		result, cycle := d.dfsForCycle(dependency.name, unvisited, currentStack, completed, cycle)
		if result {
			return result, cycle
		}

	}
	completed[project] = true
	*cycle = (*cycle)[:len(*cycle)-1]
	delete(currentStack, project)
	return false, *cycle
}

func postOrderProcessing(root *dependecyTreeNode, visited map[string]bool, result *[]string) {
	if visited[root.name] {
		return
	}
	for _, val := range root.dependencies {
		postOrderProcessing(val, visited, result)
	}
	visited[root.name] = true
	*result = append(*result, root.name)
}

func PlayDependencyTree() {
	d := InitDependencyManager()
	d.AddDependency("a", "d")
	d.AddDependency("f", "b")
	d.AddDependency("b", "d")
	d.AddDependency("f", "a")
	d.AddDependency("d", "c")

	result, cycle := d.checkCycle()
	if result{
		fmt.Println("cycle in dependencies")
		fmt.Println(cycle)
		return
	}
	r := d.GetBuildOrder()
	fmt.Println(r)
}
