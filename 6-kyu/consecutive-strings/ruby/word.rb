# frozen_string_literal: true

# Finds consecutive words
class Word
  def self.longest_consec(words, num)
    return '' if !num.positive? || num > words.count

    combine(words, num).max_by(&:length)
  end

  def self.combine(words, by)
    words.each_cons(by).map(&:join)
  end
end

