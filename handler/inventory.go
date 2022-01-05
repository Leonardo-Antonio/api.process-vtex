package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/api.process-vtex/entity"
	"github.com/Leonardo-Antonio/api.process-vtex/util"
	"github.com/Leonardo-Antonio/api.process-vtex/util/response"
	"github.com/labstack/echo/v4"
)

type Inventory struct{}

func (i *Inventory) SetBySkuAndWarehouse(c echo.Context) error {
	body := new([]entity.InventoryUpdate)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(err.Error(), nil))
	}

	if len(*body) == 0 {
		return c.JSON(http.StatusBadRequest, response.Error("ingrese por lo menos un producto a actualizar", nil))
	}

	header := util.GetHeaders(c)

	for _, v := range *body {
		go i.Update(v.SkuId, v.WarehousesId, &header, v)
		/* 	if err := i.Update(v.SkuId, v.WarehousesId, &header, v); err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(err.Error(), nil))
		} */
	}

	return c.JSON(http.StatusOK, response.Success("ok", body))
}

func (i *Inventory) Update(_skuId, _warehousesId string, header *entity.Header, inventory entity.InventoryUpdate) error {

	type body struct {
		UnlimitedQuantity bool `json:"unlimitedQuantity"`
		Quantity          int  `json:"quantity"`
	}

	data := &body{UnlimitedQuantity: inventory.Limited, Quantity: inventory.Quantity}
	buf, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf(
			"https://%s.vtexcommercestable.com.br/api/logistics/pvt/inventory/skus/%s/warehouses/%s",
			header.Store,
			_skuId,
			_warehousesId,
		),
		bytes.NewBuffer(buf),
	)
	if err != nil {
		return err
	}

	request.Header.Set("X-VTEX-API-AppKey", header.Key)
	request.Header.Set("X-VTEX-API-AppToken", header.Token)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}

	log.Println("agregado correctamente")
	return nil
}
