FROM ruby:stretch

WORKDIR /app
COPY . /app

RUN bundle install

EXPOSE 9292

CMD ["rerun", "rackup"]
