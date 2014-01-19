require 'time_series'

module Tdx
  module Data
    class Feed < TimeSeries
      def initialize(time_step = 240)
        @time_step = time_step
        super()
      end
    end
  end
end
