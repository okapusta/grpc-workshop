require 'bundler'

Bundler.require

class Application < Sinatra::Base
  get '/' do
    puts 'Hello world!'
  end
end
