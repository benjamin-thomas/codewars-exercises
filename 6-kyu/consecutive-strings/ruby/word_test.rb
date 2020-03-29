#!/usr/bin/env ruby
# frozen_string_literal: true

require 'minitest/autorun'

require_relative 'word'

# noinspection SpellCheckingInspection
class WordTest < Minitest::Test
  def test_combine_by_2
    words = %w[head shoulder knees toes]
    want = %w[headshoulder shoulderknees kneestoes]
    assert_equal(want, Word.combine(words, 2))
  end

  def test_combine_by_1
    words = %w[head shoulder knees toes]
    want = words
    assert_equal(want, Word.combine(words, 1))
  end

  def test_combine_by_3_even
    words = %w[head shoulder knees toes]
    want = %w[headshoulderknees shoulderkneestoes]
    assert_equal(want, Word.combine(words, 3))
  end

  def test_combine_by_3_odd
    words = %w[head shoulder knees toes feet]
    want = %w[headshoulderknees shoulderkneestoes kneestoesfeet]
    assert_equal(want, Word.combine(words, 3))
  end

  def test_1
    words = %w[zone abigail theta form libe zas]
    want = 'abigailtheta'
    assert_longest_consec(want, words, 2)
  end

  def test_2
    words = %w[ejjjjmmtthh zxxuueeg aanlljrrrxx
               dqqqaaabbb oocccffuucccjjjkkkjyyyeehh]
    want = 'oocccffuucccjjjkkkjyyyeehh'
    assert_longest_consec(want, words, 1)
  end

  def test_3
    # rubocop:disable Layout/LineLength
    want = 'wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck'
    # rubocop:enable Layout/LineLength
    words = %w[itvayloxrp wkppqsztdkmvcuwvereiupccauycnjutlv
               vweqilsfytihvrzlaodfixoyxvyuyvgpck]
    assert_longest_consec(want, words, 2)
  end

  def test_negative_index
    words = %w[zone abigail theta form libe zas]
    assert_longest_consec('', words, -2)
  end

  def test_high_index
    words = %w[it wkppv ixoyx 3452 zzzzzzzzzzzz]
    assert_longest_consec('', words, 15)
  end

  def test_index_zero
    words = %w[it wkppv ixoyx 3452 zzzzzzzzzzzz]
    assert_longest_consec('', words, 0)
  end

  def test_remaiming
    # rubocop:disable Layout/LineLength, Style/StringLiterals
    testing(Word.longest_consec(%w[wlwsasphmxx owiaxujylentrklctozmymu wpgozvxxiu], 2), "wlwsasphmxxowiaxujylentrklctozmymu")
    testing(Word.longest_consec(%w[it wkppv ixoyx 3452 zzzzzzzzzzzz], 3), "ixoyx3452zzzzzzzzzzzz")
    # rubocop:enable Layout/LineLength, Style/StringLiterals
  end

  def test_empty
    assert_longest_consec('', [], 3)
  end

  private

  def assert_longest_consec(want, words, by)
    assert_equal(want, Word.longest_consec(words, by))
  end

  def testing(actual, expected)
    assert_equal(actual, expected)
  end
end
