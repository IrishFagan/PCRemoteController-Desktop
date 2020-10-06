import sys

import asyncio
import websockets

async def hello(websocket, path):
	name = await websocket.recv()
	print("HELLO")

server = websockets.serve(hello, port=8080)

asyncio.get_event_loop().run_until_complete(server)
asyncio.get_event_loop().run_forever()