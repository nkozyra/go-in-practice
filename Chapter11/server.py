import concurrent.futures as futures
import grpc

import chat_pb2_grpc
import chat_pb2

class ChatServiceServicer(chat_pb2_grpc.ChatServiceServicer):
  def __init__(self, *args, **kwargs):
    self.comments = []
    pass

  def RouteComments(self, request, context):
    response = { 'commentLength': len(request.text), 'previousCommentCount': len(self.comments) }
    self.comments.append(request)
    return chat_pb2.CommentMeta(**response)

def main():
  print("starting server")
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  chat_pb2_grpc.add_ChatServiceServicer_to_server(ChatServiceServicer(), server)
  server.add_insecure_port("localhost:50051")
  server.start()
  server.wait_for_termination()

if __name__ == '__main__':
  main()