package pkg


import (
        "fmt"
        "context"
	"strings"

        "github.com/arangodb/go-driver/http"
        driver "github.com/arangodb/go-driver"

)

//Structure to pair to Json field of Jalapeno DB
type Sid struct {
        Srv6Sid string `json:"srv6sid"`
}


func ArangoDBQuery(ctx context.Context, src string, dst string, query string) (error, string) {

	//Reaching out the Jalapeno DB
        conndb, err := http.NewConnection(http.ConnectionConfig{
                Endpoints: []string{"http://10.200.99.27:32748/"},
        })
        if err != nil {
                return err, ""
        }
	//Authenticating to the Jalapeno DB
        client, err := driver.NewClient(driver.ClientConfig{
                Connection: conndb,
                Authentication: driver.BasicAuthentication("root", "jalapeno"),

        })
        if err != nil {
                return err, ""
        }
	//Opening Jalapeno DB
        db, err := client.Database(ctx, "jalapeno")

        if err != nil {
                return err, ""
        }

	if query == "Shortest Path" {

		src_query := "'" + "sr_node/" + src
		src_query += "'"
		dst_query := "'" + "sr_node/" + dst
		dst_query += "'"
		//Query to return SID list for ShortestPath among Source "src input" and Destination "dst input"
        	srv6_query := "with sr_node for v, e in outbound shortest_path " + src_query + " to " + dst_query + " sr_topology return  { srv6sid: v.srv6_sid } "

		cursor, err := db.Query(ctx, srv6_query, nil)
        	if err != nil {
              		return err, ""
        	}
        	defer cursor.Close()
        	//Slice of SIDs
		docs := make([]Sid,0)

		for {
              		var doc Sid
              		meta, err := cursor.ReadDocument(ctx, &doc)
              		_ = meta
			if driver.IsNoMoreDocuments(err) {
              			break
              		} else if err != nil {
                      		return err, ""
              		}
                docs = append(docs,doc)
      		}
		//inserire codice per creazione uSID
		var uSID string
		var split[] string
		for i, s :=range docs {
              		if i == 0 {
	      			split := strings.Split(s.Srv6Sid, ":")
				uSID = split[0] + ":" + split[1]
			}
			fmt.Printf("Sid: %s %T\n", s.Srv6Sid, s.Srv6Sid)
			split = strings.Split(s.Srv6Sid, ":")
			uSID += ":" + split[2]
      		}
		fmt.Printf("uSID: %s", uSID)
	return nil, uSID
	} else {
		//err = "No query to the Arango DB. Please double check query input"
		return nil, ""
	}
}


