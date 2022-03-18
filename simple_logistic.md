# Commerce System

## Tables

### Product
{id:1,sku:A,stock:15}
{id:2,sku:B,stock:6}

### Order
{id:1,code:b1-1,buyer:buyer1@gmail.com,items:
	[
		{sku:A,qty:1},
		{sku:B,qty:1}
	],
}

# Fulfillment System

## Tables

### Warehouse
{id:1,code:JKT}
{id:2,code:SBY}
{id:3,code:MKS}

### Product
{id:1,sku:A}
{id:2,sku:B}

### Inventory
{id:1,sku:A,qty:5,expiry_date:2022,warehouse_id:1}
{id:2,sku:A,qty:4,expiry_date:2023,warehouse_id:1}
{id:3,sku:B,qty:6,expiry_date:2023,warehouse_id:1}
{id:4,sku:A,qty:6,expiry_date:2023,warehouse_id:2}


### Outbound
{id:1,code:o1-1,order:b1-1,status:pending,items:
	[
		{sku:A,qty:1},
		{sku:B,qty:1}
	]
}


