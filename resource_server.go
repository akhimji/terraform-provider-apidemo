package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			// "address": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Required: false,
			// },

			"apiid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"firstname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"lastname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	type Payload struct {
		ID        string `json:"ID"`
		Firstname string `json:"Firstname"`
		Lastname  string `json:"lastname"`
	}
	data := Payload{
		ID:        d.Get("apiid").(string),
		Firstname: d.Get("firstname").(string),
		Lastname:  d.Get("lastname").(string),
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)
	baseurl := "http://localhost:12345/people/"
	url := baseurl + d.Get("apiid").(string)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	d.SetId(d.Get("apiid").(string))
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {

	type Payload struct {
		ID        string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
	data := new(Payload)

	baseurl := "http://localhost:12345/people/"
	url := baseurl + d.Get("apiid").(string)

	resp, err := http.Get(url)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	d.Set("apiid", data.ID)
	d.Set("firstname", data.Firstname)
	d.Set("lastname", data.Lastname)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	d.Partial(true)

	type Payload struct {
		ID        string `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	data := Payload{
		ID:        d.Get("apiid").(string),
		Firstname: d.Get("firstname").(string),
		Lastname:  d.Get("lastname").(string),
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)
	baseurl := "http://localhost:12345/people/"
	url := baseurl + d.Get("apiid").(string)
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	d.Partial(false)

	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	baseurl := "http://localhost:12345/people/"
	url := baseurl + d.Get("apiid").(string)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	d.SetId("")
	return nil
}
