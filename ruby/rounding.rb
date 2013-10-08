#!ruby

def create_charges(count = 20)
  count.times.map { (-1000+Random.rand(2000)) }
end

def assess(charges, fee = 0.1)
  expect = charges.inject(:+) / fee

  remainder = 0
  actual = 0
  sum = 0
  charges.each do |chrg|
    hold_amount_cents = chrg * fee
    hold_amount = (hold_amount_cents + remainder).round
    remainder += hold_amount_cents - hold_amount
    
    puts("chrg=%4d¢ wanted=%5.2f¢ got=%6.2f¢ remainder=%5.2f¢ [sum=%6.0f¢]" % [chrg, hold_amount_cents, hold_amount, remainder, sum])
    
    sum += hold_amount
  end
  
  remainder
end

def simulate()
  charges = create_charges(20) 
  del = assess(charges, 0.1)
  if del.abs > 1
    puts("%s del=%f" % [charges.join(' '), del])
  end
end

100_000.times do
  simulate()
end

# simulate()
# assess([3066,4052])