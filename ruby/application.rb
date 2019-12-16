require 'bundler'
require "sinatra/json"

Bundler.require

class Application < Sinatra::Base
  get '/' do
    json kiszka: 'majonez'
  end
end
