package main

import "fmt"
import "math/rand"
import "time"

type packet struct {
    srce int
    dest int
}

func make_bundle(width int) (bs []chan packet) {
    bs = make([]chan packet,width)
    for i := 0; i < width; i++ {
        bs[i] = make(chan packet)
    }
    return
}

func make_injector(chs []chan packet) {
    n := len(chs)
    go func() {
        for {
            s := rand.Intn(n)
            t := rand.Intn(n)
            fmt.Printf("Injection of packet {%v --> %v}.\n", s, t)
            chs[s] <- packet{s,t}
            time.Sleep(time.Duration(2)*time.Second)
        }
    } ()
}
func which_neighbor(current int, dest int) uint{
    var n uint = 0
    for current%2 == dest%2{
        current = current/2
        dest = dest/2
        n = n + 1
    }
    return n
}


func node(index int, in chan packet, out [](chan packet), report chan packet) {
    for {
        p := <-in
        if p.dest == index {
            fmt.Printf("Node %v received packet from %v.\n",index,p.srce)
            report <- p
        } else {
            var next_neighbor_bit uint = which_neighbor(index, p.dest)
            var next_neighbor int = index^(1<<next_neighbor_bit)
            fmt.Printf("Node %v forwarding packet {%v --> %v}.\n",index,p.srce,p.dest)
            out[next_neighbor] <- p
        }
    }
}

func make_cube(chs []chan packet, rep chan packet) {
    n := len(chs)
    i := 0
    for i < n {
        go node(i, chs[i], chs , rep)
        i = i + 1
    }
}
func pow_two(a int) int{
    if a==0{
        return 1
    }else if a%2==1{
        return pow_two(a-1)*2
    }else{
        return pow_two(a/2)*pow_two(a/2)
    }
}

func main() {
    fmt.Println("Enter integer for number of nodes in the hypercube: ")
	var n int    
    fmt.Scanln(&n)
    n = pow_two(n)
    chs := make_bundle(n)
    rep := make(chan packet)
    make_injector(chs)
    make_cube(chs,rep)
    for x := range rep {
        fmt.Printf("Packet {%v --> %v} routed.\n", x.srce, x.dest)
    }
}