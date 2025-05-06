# GOBMS
package main
import (
    "fmt"
    "sync"
    "time"
    "context"
    )
func Consumer(id int, taskQueue chan int, wg *sync.WaitGroup, responseMap map[int]string, ctx context.Context){
    defer wg.Done()
    
    for {
        select {
            case <-ctx.Done():
               fmt.Println("Process is completed for goroutine id",id)
            case val, ok := <-taskQueue:
                if !ok{
                    fmt.Println("TaskQueue is empty now")
                    return
                }
                    fmt.Println("Data is received by gorotuine id",id+1," with value as ",val)
                    api1 := "x"
                    api2 := "y"
                    api3 := "z"
                    if val%2==0 {
                        responseMap[val]= "true"+" "+api1+" "+api2 +" "+api3
                    }else{
                        responseMap[val]= "true"+ " " + api1+ " " +api2+ " " +api3
                    }
                    time.Sleep(1*time.Second)
        }
    }
    // for val := range taskQueue{
    //     fmt.Println("Data is received by gorotuine id",id+1," with value as ",val)
    //     if val%2==0 {
    //         responseMap[val]=true
    //     }else{
    //         responseMap[val]=false
    //     }
        // api1 -resp
        //api2 - 
        //api3
        //true+ api1 + api2 + api3
        // time.Sleep(2*time.Second)
    // }
}
func main() {
  slice := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
  taskQueue := make(chan int, 5)
  responseMap := make(map[int]string,30)
  ctx, cancel := context.WithTimeout(context.Background(),3*time.Second)
  defer cancel()
  var wg sync.WaitGroup
  const maxGo int =5
  for i:=0; i<maxGo; i++ {
      wg.Add(1)
      go Consumer(i,taskQueue,&wg,responseMap,ctx)
  }
  wg.Add(1)
  go func(){
      defer wg.Done()
      for i:=0;i <len(slice); i++{
          taskQueue <- slice[i]
          fmt.Println("Data pushed successfully ",slice[i])
      }
      close(taskQueue)
  }()
  wg.Wait()
  for key,value :=range responseMap{
      fmt.Println("key ->",key,"value ->",value)
  }
  fmt.Println("Task Completed!")
}



