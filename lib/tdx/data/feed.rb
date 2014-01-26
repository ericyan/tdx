require 'time_series'

module Tdx
  module Data
    class Feed < TimeSeries
      def initialize(time_step = 240, *args)
        @time_step = time_step
        super(*args)
      end

      def timestamps
        Hash[@data_points.sort].keys
      end

      def extract(element)
        timestamps = @data_points.keys
        elements = @data_points.values.collect { |dp| dp.data.fetch(element) }

        Data::Feed.new(@time_step, timestamps, elements)
      end

      def slice(timeframe)
        timeframe = timeframe.inject({}) do |timeframe, (key, value)|
          timeframe[key] = (value.kind_of? Time) ? value : value.to_eod_time
          timeframe
        end

        Data::Feed.new(@time_step, super.to_a)
      end

      def scale(new_time_step)
        if new_time_step % @time_step == 0
          scaled_data_points = []

          @data_points.values.each_slice(new_time_step / @time_step) do |slice|
            scaled_data_points << Data::Candlestick.new(
              slice.last.timestamp,
              {
                open: slice.first.data[:open],
                high: slice.max_by { |c| c.data[:high] }.data[:high],
                low: slice.min_by { |c| c.data[:low] }.data[:low],
                close: slice.last.data[:close],
                turnover: slice.inject(0) { |sum, c| sum + c.data[:turnover] },
                volume: slice.inject(0) { |sum, c| sum + c.data[:volume] }
              }
            )
          end

          Data::Feed.new(new_time_step, scaled_data_points)
        else
          raise ArgumentError, 'New time step must be an integral multiple of the old one'
        end
      end
    end
  end
end
