package db

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/vpalahia/Cache/types"
)

func FetchFromDB(conn *amqp.Connection) error {
	var todoList []types.Todo

	fetchQuery := `SELECT id, task, description FROM todotasks`
	err := db.Select(&todoList, fetchQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}

	byteList, err := json.Marshal(todoList)
	if err != nil {

		return err
	}

	err = publish(conn, "cachedata", byteList)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//publish to queue
func publish(conn *amqp.Connection, q string, msg []byte) error {

	// create a channel through which we publish
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// create the payload with the message that we specify in the arguments
	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
	}

	// publish the message to the queue specified in the arguments
	if err := ch.Publish("", q, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] failed to publish to queue %v", err)
	}
	return nil

}
