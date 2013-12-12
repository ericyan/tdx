require 'bindata'
require 'date'

module Tdx
  class Parser < BinData::Record
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
      super(data)

      return {
        date:     Date.parse(date.to_s),
        open:     open / 100.00,
        high:     high / 100.00,
        low:      low / 100.00,
        close:    close / 100.00,
        turnover: turnover.to_i,
        volume:   volume
      }
    end
  end
end
