module Tdx
  module Data
    class File < File
      def each_record(record_size = 32, &block)
        (self.size / record_size).times do |sn|
          self.pos = sn * record_size
          yield self
        end
      end
    end
  end
end
