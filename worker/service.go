package worker

import (
	"encoding/json"
	"log"

	"github.com/workflow-interoperability/manufacturer/lib"
	"github.com/workflow-interoperability/manufacturer/types"

	"github.com/gorilla/websocket"
)

func publishPIIS(piisid string, IM *types.IM, sub string, conn *websocket.Conn) (bool, error) {
	// get piis
	processData, err := lib.GetPIIS("http://127.0.0.1:3000/api/PIIS/" + piisid)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if !(processData.To.ProcessID == IM.Payload.WorkflowRelevantData.From.ProcessID && processData.To.ProcessInstanceID == IM.Payload.WorkflowRelevantData.From.ProcessInstanceID && processData.To.IESMID == IM.Payload.WorkflowRelevantData.From.IESMID) {
		return false, nil
	}
	IM.Payload.WorkflowRelevantData.To.ProcessInstanceID = processData.From.ProcessInstanceID
	// create piis
	id := lib.GenerateXID()
	newPIIS := types.PIIS{
		ID:   id,
		From: IM.Payload.WorkflowRelevantData.From,
		To:   processData.From,
		IMID: IM.ID,
		SubscriberInformation: types.SubscriberInformation{
			Roles: []string{},
			ID:    sub,
		},
	}
	pPIIS := types.PublishPIIS{newPIIS}
	body, err := json.Marshal(&pPIIS)
	if err != nil {
		log.Println(err)
		return false, err
	}
	err = lib.BlockchainTransaction("http://127.0.0.1:3000/api/PublishPIIS", string(body))
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
