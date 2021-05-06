package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SevereCloud/vksdk/v2/api"
	giftsobj "github.com/SevereCloud/vksdk/v2/object"
)

type ProfileInfo struct {
	FollowersCount           int
	Followers                []int
	SubscriptionsUsersCount  int
	SubscriptionsUsers       []int
	SubscriptionsGroupsCount int
	SubscriptionsGroups      []int
	GiftsCount               int
	Gifts                    []giftsobj.GiftsGift
}

func main() {
	ids := os.Args[1:len(os.Args)]
	token := os.Getenv("token")
	vk := api.NewVK(token)
	for _, id := range ids {
		int_id, _ := strconv.Atoi(id)
		info := GetInfo(int_id, vk)
		PrintInfo(info)
	}
}

func GetInfo(id int, vk *api.VK) ProfileInfo {
	followers, _ := vk.UsersGetFollowers(api.Params{"user_id": id})
	subscriptions, _ := vk.UsersGetSubscriptions(api.Params{"user_id": id})
	gifts, err := vk.GiftsGet(api.Params{"user_id": id})
	if err != nil {
		gifts = api.GiftsGetResponse{0, nil}
	}
	return ProfileInfo{
		followers.Count,
		followers.Items,
		subscriptions.Users.Count,
		subscriptions.Users.Items,
		subscriptions.Groups.Count,
		subscriptions.Groups.Items,
		gifts.Count,
		gifts.Items,
	}
}

func PrintInfo(info ProfileInfo) {
	fmt.Printf("%d FOLLOWERS count with ids %v\n", info.FollowersCount, info.Followers)
	fmt.Printf("%d SUBSCRIPTIONS count (%d users, %d groups)\n", info.SubscriptionsUsersCount + info.SubscriptionsGroupsCount, info.SubscriptionsUsersCount, info.SubscriptionsGroupsCount)
	fmt.Println()
	fmt.Printf("Subscriptions USERS %v\n", info.SubscriptionsUsers)
	fmt.Printf("Subscriptions GROUPS %v\n", info.SubscriptionsGroups)
	fmt.Println()
	fmt.Printf("%d Gifts count \n", info.GiftsCount)
}
