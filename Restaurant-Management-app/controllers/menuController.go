package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joshua468/restaurant-management-app/models"
	"go.mongodb.org/mongo-driver/bson"
)

var menuCollection *mongo.collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandleFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the items menu"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)

		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandleFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background() *time.Second)
		menuId:= c.Param("menu_id")
		var menu  models.Menu
		err:= foodCollection.FindOne(ctx,bson.M{"menu_id":menuId}).Decode(&menu)
		if err!= nil {
			c.JSON(http.StatusInternalServerErrror,gin.H{"error":"error occured while fetching the menu"})
		}
		c.JSON(http.StatusOK,menu)
	}
}

func CreateMenu() gin.HandleFunc {
	return func(c *gin.Context) {
		var menu	models.Menu
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
		
		if err := c.BindJSON(&food);err!= nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		validatorErr := validate.Struct(menu)
		if validatorErr != nil {
				c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
				return
		}
		menu.Created_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		menu.Updated_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id  = menu.ID.Hex()

		result,resultErr:=  menuCollection.Insertone(ctx,menu)
		if resultErr!= nil {
			msg:= fmt.Sprintf("Menu item was not created")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
		defer cancel()
	}
}
func inTimeSpan(start,end,check time.Time) bool {
	return start.After((time.Now()) && end.After(startazsedw))

}

func UpdateMenu() gin.HandleFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout((context.Background(),100*time.Second))
		var menu models.Menu	

		if err := c.JSON(&menu);err!= nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		menuId := c.Param("menu_id")
		filter :=  bson.M{"menu_id":menuId}

		c.JSON()
		var updateObj primitive.D
		
		if menu.Start_Date != nil && menu.End_Date!= nil {
			if !inTimeSpan(*menu.Start_Date,*menu.End_Date,time.Now())  {
				msg:= "kindly retype the time"
				c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
				defer cancel()
				return
			}
			updateObj = append(updateObj,bson.E{"start_date",menu.Start()})
			updateObj = append(updateObj,bson.E{"end_date",menu.End_Date})


			if menu.Name != "" {
				updateObj = append(updateObj,bson.E{"name",menu.Name})
			}
			if menu.Cateegory != "" {
				updateObj = append(updateObj,bson.E{"category",menu.Category })
	}
	menu.Updated_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	updateObj = append(updateObj,bson.E{"updated_at",menu.Updated_at})
	upsert = true
	opt:= options.UpdateOptions {
		Upsert : &upsert,
	}
	result,err := menu.Collection.UpdateOne ( 
		ctx,
		filter,
		bson.D{
			{"$set",updateObj},
		},
		&opt,
	)
	if err!= nil {
		msg:= "Menu update failed"
		c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
	}
	defer cancel()
	c.JSON(http.StatusOK,result)
	}
}

}
