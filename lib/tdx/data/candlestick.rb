require 'time_series'

module Tdx
  module Data
    class Candlestick < DataPoint
      def initialize(timestamp, data)
        raise ArgumentError, 'Invalid candlestick data' unless valid?(data)
        super
      end

      private
        def valid?(data)
          validity = true

          [:open, :high, :low, :close].each do |key|
            validity = validity and data.has_key?(key)
            validity = validity and data[key].kind_of? Numeric
          end

          return validity
        end
    end
  end
end
