require 'time_series'

module Tdx
  module Data
    class Feed < TimeSeries
      def initialize(time_step = 240, *args)
        @time_step = time_step
        super(*args)
      end

      def slice(timeframe)
        Data::Feed.new(@time_step, super.data_points)
      end
    end
  end
end
