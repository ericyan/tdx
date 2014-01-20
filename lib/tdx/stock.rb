module Tdx
  class Stock
    def initialize(symbol)
      @stock_code, @exchange = parse_symbol(symbol)
    end

    def quotes(time_step = 240)
      case time_step
      when 5
        file = Data::File.open("data/#{@exchange.to_s.downcase}#{@stock_code}.5", 'rb')
        @quotes = Parsers::FiveMinutes.parse(file)
      when 240
        file = Data::File.open("data/#{@exchange.to_s.downcase}#{@stock_code}.day", 'rb')
        @quotes = Parsers::EoD.parse(file)
      else
        raise ArgumentError, 'Invalid time step'
      end
    end

    private
      def parse_symbol(symbol)
        stock_code, exchange = symbol.split('.')
        exchange = exchange.to_s.upcase.to_sym
        unless stock_code =~ /^\d{6}$/ and [:SZ, :SH].include? exchange
          raise ArgumentError, 'Invalid stock symbol'
        end

        return [stock_code, exchange]
      end
  end
end
