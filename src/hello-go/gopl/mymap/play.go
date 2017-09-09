package mymap

import (
	"hello-go/gopl"
	"fmt"
	 _ "sort"
//	"bufio"
//	"os"
)

func init()  {
	gopl.AddRunFunc(run)
}

func run()  {
	fmt.Println( "ok  i am playing with the map type ")

	ages := make(map[string]int)
	fmt.Printf("type is %T \n", ages )
	ages["alice"] = 34
	ages["chalie"] = 34
	fmt.Println(ages)


	// 初始化
	func(){
		// 字面值创建
		ages := map[string]int {
			"alice":34,
			"chalie": 34,
		}
		fmt.Println(ages)
	}()

	// 空 map
	func(){
		emptyMap := map[string]int{}
		fmt.Println(emptyMap)
	}()

	// 读 访问
	fmt.Println(ages["alice"])
	// 删除
	delete(ages, "alice")
	fmt.Println(ages)

	// 访问不存在时 返回值类型的零值
	fmt.Println(ages["notExistsKey"])

	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)

	// 简短赋值语法 |  复合语句
	ages["bob"] += 1
	fmt.Println(ages["bob"])
	// 递增
	ages["bob"] ++
	fmt.Println(ages["bob"])

	// 遍历
	for name, age := range ages{
		fmt.Printf("%s\t%d \n", name, age)
	}

	// 按key排序
	func(){
		/*
		names := make([]string, 0, len(ages))
		// var names []string
		// 忽略map的value部分
		for name := range ages {
			names = append(names, name)
		}
		sort.Sort(names)

		// 忽略map的key部分
		for _, name := range names{
			fmt.Printf("%s\t%d", name, ages[name])
		}
		*/
	}()

	func(){
		// 空
		var ages map[string]int
		fmt.Println(ages == nil)
		fmt.Println(len(ages) == 0)

		age , ok := ages["bob"]
		if !ok{
			fmt.Println(age)
		}
		// ages["myKey"] = 1 // 注意写之前一定要保证

	}()

        /*
	fmt.Println("play with the graph ！")
	// 玩玩图
	addEdge("node-1","node-2")
	addEdge("node-1","node-3")
	addEdge("node-2","node-3")

	fmt.Printf("%b \t %v " , hasEdge("node-1","node-2"))

	// 模拟set
	func(){
		seen := make(map[string]bool)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan(){
			line := input.Text()
			if !seen[line]{
				seen[line] = true
				fmt.Println(line)
			}
		}

		if err := input.Err(); err != nil{
			fmt.Fprintf(os.Stderr,"dedup %v\n", err)
			os.Exit(1)
		}
	} /*()  // 暂时不运行*/



}

var m = make(map[string]int)
func k(list []string) string {return fmt.Sprintf("%q",list)}
func Add(list []string) { m[k(list)] ++ }
func Count(list []string) int  {
	return m[k(list)]
}


var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}