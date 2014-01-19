require 'bindata'
require 'date'

module Tdx
  module Parsers
    class FiveMinutes
      def self.parse(file)
        quotes = Tdx::Data::Feed.new(5)

        (file.size / 32).times do |line|
          file.pos = line * 32
          quotes << Quote.new.read(file)
        end

        return quotes
      end

      class Quote < BinData::Record
        endian  :little

        uint16  :date       # Date, as (yyyy - 2004) * 2048 + mmdd
        uint16  :time       # Minutes since 00:00:00 (UTC+8)
        uint32  :open       # Opening price, in cents
        uint32  :high       # Highiest price, in cents
        uint32  :low        # Lowest price, in cents
        uint32  :close      # Closing price, in cents
        float   :turnover   # Turnover
        uint32  :volume     # Volume
        skip    length: 4   # Reserved

        def read(data)
          super

          return Tdx::Data::Candlestick.new(
            Time.new(date / 2048 + 2004, (date % 2048) / 100, (date % 2048) % 100,time / 60, time % 60, 0, '+08:00'),
            {
              open:     open / 100.00,
              high:     high / 100.00,
              low:      low / 100.00,
              close:    close / 100.00,
              turnover: turnover.to_i,
              volume:   volume
            }
          )
        end
      end
    end
  end
end
