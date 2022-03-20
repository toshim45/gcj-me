# Commerce System

## Tables

### Seller
{id:1,name:TokoSatu}
{id:2,name:TokoDua}

### Product
{id:1,sku:A,stock:15,seller_id:1}
{id:2,sku:B,stock:6,seller_id:2}

### Order
{id:1,code:b1-1,buyer:buyer1@gmail.com,address:"jakarta",items:
	[
		{sku:A,qty:1},
		{sku:B,qty:1}
	],
}

# Fulfillment System


## Overview
Our fulfillment system has 3 warehouses, jakarta, surabaya And makassar. Also we have the same product data with commerce system. Each warehouse will manage its inbound, outbound and inventories. When order come from commerce system, it will create outbound record at fulfillment system, then allocate inventory based on nearest expiry date, longest inbound date and nearest warehouse.

## Tables

### Warehouse
{id:1,code:JKT}
{id:2,code:SBY}
{id:3,code:MKS}

### Product
{id:1,sku:A}
{id:2,sku:B}

### Inventory
{id:1,sku:A,qty:5,expiry_date:2022,warehouse_id:1,created:2021}
{id:2,sku:A,qty:4,expiry_date:2023,warehouse_id:1,created:2021}
{id:3,sku:B,qty:6,expiry_date:2023,warehouse_id:1,created:2022}
{id:4,sku:A,qty:6,expiry_date:2023,warehouse_id:2,created:2021}

### Outbound Status
[pending, allocated, picking, consolidation, checker, dispatch]

### Outbound
{id:1,code:o1-1,order:b1-1,status:pending,address:"jakarta",items:
	[
		{sku:A,qty:1},
		{sku:B,qty:1}
	]
}


# System Design

## buyer journey

buyer --> Commerce   -->  Fulfillment --> (3pl) courier --> buyer

## flow Fulfillment
- product read only, redundancy (synchronization issue)
- incoming outbound (throttle issue)
- inventory allocation -> select * from inventory where sku = <> order by expiry_date asc (panel2)
	- inventory size (large data, read/write 50%, data issue)
- request pickup (integration issue, retry)
- dashboard to show product which inventory almost running out & expired
- how do you know your system is OK ?
