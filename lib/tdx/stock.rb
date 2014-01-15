module Tdx
  class Stock
    attr_reader :quotes

    def initialize(symbol, interval = nil)
      if interval == 5
        file = File.open("data/#{symbol}.5", 'rb')
        @quotes = Parsers::FiveMinutes.parse(file)
      else
        file = File.open("data/#{symbol}.day", 'rb')
        @quotes = Parsers::EoD.parse(file)
      end
    end
  end
end
