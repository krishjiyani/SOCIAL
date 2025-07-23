		package db

import (
	"context"
	"fmt"
	"log"
	"database/sql"

	"krishjiyani/SOCIAL/internal/store"
	"math/rand"
)
var userNames = []string{
    "alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy",
    "karl", "laura", "mike", "nancy", "oliver", "paul", "quinn", "rachel", "steve", "tina",
    "ursula", "victor", "wanda", "xavier", "yasmine", "zack", "adam", "bella", "carter", "diana",
    "edward", "fiona", "george", "hannah", "isaac", "jasmine", "kevin", "linda", "matt", "nora",
    "oscar", "penny", "quentin", "ryan", "sophia", "travis", "uma", "vincent", "willow", "zoe",
}

var titles = []string{ "Unlock Your Potential",
"Embrace the Journey", "Find Your Balance","Create Lasting Change","Discover Your Passion",
"Live with Intention","Nurture Your Mind","Explore New Horizons","Cultivate Positive Habits",
"Transform Your Space","Connect with Nature","Master Your Craft","Fuel Your Creativity",
"Achieve Your Goals","Celebrate Small Wins","Invest in Yourself","Practice Daily Gratitude",
"Overcome Your Fears","Design Your Life","Inspire and Empower",}

var contents = []string{
	"Discover how to unlock your true potential with simple daily practices.",
	"Join us on a journey of self-discovery and personal growth.",
	"Learn effective strategies to find balance in your busy life.",
	"Explore actionable steps to create lasting change in your habits.",
	"Uncover your passion and turn it into a fulfilling career.",
	"Understand the importance of living with intention and purpose.",
	"Nurture your mind with mindfulness techniques and meditation.",
	"Find out how to explore new horizons and expand your experiences.",
	"Cultivate positive habits that lead to a happier, healthier life.",
	"Transform your living space into a sanctuary of peace and creativity.",
	"Connect with nature and its benefits for mental well-being.",
	"Master your craft with tips from industry experts and enthusiasts.",
	"Fuel your creativity with exercises and inspiration from various fields.",
	"Achieve your goals by setting clear, actionable steps and milestones.",
	"Celebrate small wins to boost your motivation and confidence.",
	"Invest in yourself through continuous learning and self-improvement.",
	"Practice daily gratitude to enhance your overall happiness.",
	"Overcome your fears by facing them head-on with practical techniques.",
	"Design your life around your values and aspirations for fulfillment.",
	"Inspire and empower others by sharing your journey and insights.",
	"Embrace change as a natural part of life and personal growth.",
}

var tags = []string{
	"Personal Development","Mindfulness","Productivity","Wellness","Creativity",
	"Self-Improvement","Motivation","Lifestyle","Mental Health",
	"Inspiration","Goal Setting","Sustainability","Travel","Healthy Living","Work-Life Balance",
	"Happiness","Entrepreneurship","Fitness","Relationships","Learning",
}
var comments = []string{
	"I completely agree with your thoughts.",
	"Thanks for the tips, very helpful.",
	"Interesting perspective, I hadn't considered that.",
	"Thanks for sharing your experience.",
	"Well written, I enjoyed reading this.",
	"This is very insightful, thanks for posting.",
	"Great advice, I'll definitely try that.",
	"I love this, very inspirational.",
	"Thanks for the information, very useful.",
}

func Seed(store store.Storage, db *sql.DB)  {
	ctx := context.Background()

	users := generateUsers(100)
	tx,_ := db.BeginTx(ctx,nil)

for _, user := range users{
	if err := store.Users.Create(ctx, user); err != nil{
	_ = tx.Rollback()
	log.Println("error creating user:",err)
	return
	}
 }
 tx.Commit()

posts := generatePosts(200, users)
for _, post := range posts {
	if err := store.Posts.Create(ctx, post); err != nil {
		log.Println("Error Creating post:", err)
		return
	}
}
comments := generateComments(500, users, posts)
for _, comment := range comments {
	if err := store.Comments.Create(ctx, comment); err != nil {
		log.Println("Error Creating post:", err)
		return
	}
}
	log.Println("Seeding complete")
}

func generateUsers(num int) []* store.User {
	users := make([] *store.User, num)

	for i:= 0; i<num;i++{
		users[i] = &store.User{
			Username:userNames[i%len(userNames)]+fmt.Sprintf("%d", i),
			Email:userNames[i%len(userNames)]+fmt.Sprintf("%d", i) + "@example.com",
			//Role: store.Role{
		// 		Name: "user",
		// },
	}
}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
	user := users[rand.Intn(len(users))]

posts[i] = &store.Post{
	UserID: user.ID,
	Title: titles[rand.Intn(len(titles))],
	Content:titles[rand.Intn(len(contents))],
	Tags: []string{
		tags[rand.Intn(len(titles))],
		tags[rand.Intn(len(titles))],
},
}
	}
	return posts
}
func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}