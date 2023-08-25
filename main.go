package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const MAX_FAMILY = 2
	const MAX_PASSENGER = 4

	fmt.Print("Input the number of families : ")
	var numberOfFamily int
	fmt.Scanln(&numberOfFamily)

	fmt.Print("Input the number of members in the family (separated by a space) :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	array := strings.Split(scanner.Text(), " ")

	arrayLen := len(array)
	if arrayLen != numberOfFamily {
		log.Fatal("Input must be equal with count of family")
	}

	numberOfMember := []int{}
	for i := 0; i < arrayLen; i++ {
		temp, err := strconv.Atoi(array[i])
		numberOfMember = append(numberOfMember, temp)
		if err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(numberOfMember)
	family := MAX_FAMILY
	passenger := MAX_PASSENGER
	numberOfBus := 0

	for numberOfMember[len(numberOfMember)-1] != 0 {
		memberOfFamily := numberOfMember[len(numberOfMember)-1]

		if memberOfFamily >= passenger {
			memberOfFamily -= passenger
			numberOfMember[len(numberOfMember)-1] = memberOfFamily
			sort.Ints(numberOfMember)
			numberOfBus++
			family = MAX_FAMILY
			passenger = MAX_PASSENGER
		} else {
			passenger -= memberOfFamily
			numberOfMember[len(numberOfMember)-1] = 0
			sort.Ints(numberOfMember)
			family--
			if family <= 0 {
				numberOfBus++
				family = MAX_FAMILY
				passenger = MAX_PASSENGER
			}
		}

		if numberOfMember[len(numberOfMember)-1] == 0 && passenger < MAX_PASSENGER {
			numberOfBus++
			family = MAX_FAMILY
			passenger = MAX_PASSENGER
		}

	}

	fmt.Println("Minimum bus required is : ", numberOfBus)
}
