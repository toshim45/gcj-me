import uuid
import random
import time
import datetime
import json
import pymysql


rows = 1000000
batch = 100

skus = ['SKU-NONKR-CHAIR', 'SKU-MOMOGI', 'SKU-HP-XIAOMU', 'SKU-BENG-BENG', 'SKU-BENG-BENG-MOMOGI']

insert_stmt = "INSERT INTO inventory (id, sku, expiry_date, created) VALUES (%s, %s, %s, %s)"

class Inventory:
	def __init__(self):
		self.id = uuid.uuid4()
		self.sku = random.choice(skus)
		self.created = time.time()
		self.expiry_date = datetime.datetime(2023, 12, 30)

	def str(self):
		return "id:{} sku:{} created:{} expiry_date:{}".format(str(self.id),self.sku,datetime.datetime.fromtimestamp(self.created), self.expiry_date.strftime('%Y-%m-%d'))

def dummy_insert(cur):
	inventory_values = []
	for i in range(batch):
		inv_obj = Inventory()
		# print("inventory {}".format(inv_obj.str()))
		inventory_values.append([str(inv_obj.id),inv_obj.sku,inv_obj.expiry_date,datetime.datetime.fromtimestamp(inv_obj.created)])

	cur.executemany(insert_stmt, inventory_values)



conn = pymysql.connect(user="root", password="root", host="localhost", database="wms_v1", autocommit=True)

batch_times = 0

if batch > rows:
	batch = rows
	batch_times = 1
else:
	batch_times = rows // batch

print("executing batches: {}".format(batch_times))

try:
	cur = conn.cursor()
	for i in range(batch_times):
		print("batch: {}".format(i+1))
		dummy_insert(cur)
except:
	raise
finally:
	cur.close()



