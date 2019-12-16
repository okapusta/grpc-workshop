require 'bundler'
require "sinatra/json"

Bundler.require

require_relative 'lib/server_messages_services_pb'

class Application < Sinatra::Base
  helpers do
    def grpc_client
      @client ||= MyAwesomeService::Stub.new(
        ENV['GRPC_HOST'], :this_channel_is_insecure
      )
    end
  end

  get '/' do
    begin
      request = GetSomeDataRequest.new(id: 1)
      resp = grpc_client.get_some_data(request)
      json resp.to_h
    rescue GRPC::Unavailable => e
      json error: 'grpc-unavailable'
    rescue => e
      json error: e
    end
  end
end
