require 'bindata'
require 'date'

module Tdx
  module Parsers
    class EoD
      def self.parse(file)
        quotes = Tdx::Data::Feed.new
        file.each_record { |data| quotes << Quote.new.read(data) }

        return quotes
      end

      class Quote < BinData::Record
        endian  :little

        uint32  :date       # Date, as yyyymmdd
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
            Date.parse(date.to_s).to_time + 3600 * 15,
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
