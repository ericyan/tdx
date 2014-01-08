module Tdx
  class Stock
    attr_reader :quotes

    def initialize(symbol, range = nil)
      data = File.open("data/#{symbol}.day", 'rb')

      @quotes = []
      (data.size / 32).times do |line|
        data.pos = line * 32
        @quotes << Parsers::EoD.read(data)
      end
    end
  end
end
