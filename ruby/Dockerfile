FROM ruby:stretch

WORKDIR /app
COPY . /app

RUN bundle install

EXPOSE 9292

CMD ["bundle", "exec", "rerun -b 'rackup -o 0.0.0.0'"]
