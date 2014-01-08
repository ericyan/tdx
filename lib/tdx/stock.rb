module Tdx
  class Stock
    attr_reader :quotes

    def initialize(symbol, range = nil)
      file = File.open("data/#{symbol}.day", 'rb')
      @quotes = Parsers::EoD.parse(file)
    end
  end
end
