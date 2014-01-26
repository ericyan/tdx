require 'date'

class Date
  def to_eod_time
    self.to_time.getlocal("+08:00") + (60 * 60 * 15)
  end
end
