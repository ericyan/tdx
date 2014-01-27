require 'spec_helper'
require 'date'

describe Tdx::Stock do
  describe "#new" do
    it "takes a stock symbol" do
      stock = Tdx::Stock.new('000300.SH')
      stock.should be_a_kind_of Tdx::Stock
    end
  end
end
