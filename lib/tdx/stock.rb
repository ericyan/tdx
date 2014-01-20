module Tdx
  class Stock
    def initialize(symbol)
      @symbol = symbol
    end

    def quotes(time_step = 240)
      case time_step
      when 5
        file = Data::File.open("data/#{@symbol}.5", 'rb')
        @quotes = Parsers::FiveMinutes.parse(file)
      when 240
        file = Data::File.open("data/#{@symbol}.day", 'rb')
        @quotes = Parsers::EoD.parse(file)
      else
        raise ArgumentError, 'Invalid time step'
      end
    end
  end
end
