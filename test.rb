COMMAND = "go run main.go"
#COMMAND = "ruby 4.rb"

Dir.glob("test/test_*").each { |test|
  STDOUT.flush
  width = test.match(%r{test/test_(\d+)})[1]
  print "Testing #{test} @ #{width}: "
  result = test.sub(%r{/test}, "/out")
  if system("#{COMMAND} #{width} < #{test} | diff #{result} -")
    puts "."
  else
    puts "FAIL"
  end
}
