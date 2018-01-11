package main
import "fmt"

type Token struct {
    data string
    recipient int
    ttl int
    globInfo bool// Сообщение для каждого о конце эмуляции, чобы закрыть горутины
}

type Sub struct{
    adress int
    inadr chan Token
    outadr chan Token
}

func (s* Sub)isRecipient(){
    msg := <-s.inadr
    for msg.globInfo ==true{
      if msg.ttl <= 0 {
        println("time out in adress", s.adress)
      } else if s.adress == msg.recipient {
        println ("Сообщение получено абонентом ",s.adress)
        //...
      } else{
        msg.ttl -=1;
        s.outadr <- msg
      }
      msg = <-s.inadr
    }
    println("закрываю абонена №",s.adress)
    s.outadr <- msg
  }

func main() {
    const k = 50
    var rec int
    fmt.Println("Напишите адресата сообщения: ")
    fmt.Scanln(&rec)
    t := Token{"some data",rec,111,true}
// инициализация
    var subs [k] Sub;
    var channels [k] chan Token
    for i:=0;i<k;i++ {
      channels[i] = make(chan Token)
    }
    for i:=0;i<k;i++ {
      if i==0 {
        subs[i].inadr = channels[k-1]
        subs[i].outadr = channels[i]
        subs[i].adress =i;
        go subs[i].isRecipient()
      } else{
        subs[i].inadr = channels[i-1]
        subs[i].outadr = channels[i]
	      subs[i].adress =i;
	      go subs[i].isRecipient()
      }}
// ...
    channels[k-1]<-t

    var input string
    fmt.Scanln(&input)
    channels[k-1]<-Token{"",-1,k+1,false}
    t= <-channels[k-1]
    println(t.data)
}
