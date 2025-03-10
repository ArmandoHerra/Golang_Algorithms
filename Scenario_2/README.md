# Scenario 2 - Palindrome Check

## Problem Description

A palindrome is a string that reads the same forward and backward. Your task is to write a function to determine if a given string is a palindrome.

Examples:

- "racecar" → True (because it reads racecar in both directions)
- "hello" → False (because backward it’s olleh, which differs from hello)
- "A" → True (single-character strings are trivially palindromic)
- "abba" → True
- "Madam" → Typically this is False if we’re doing a strict character-by-character check with case sensitivity. If we wanted case-insensitive, we’d first convert to a single case (e.g., lowercase).

## Constraints

- The input string can be empty; an empty string is generally considered a palindrome by definition.
- Decide if you want the check to be case-sensitive or case-insensitive (and specify which). The examples below assume case-sensitive by default, but you can easily extend to case-insensitive.
- Unicode considerations: if you’re dealing with Unicode characters (like emojis or accented characters), you may want to treat them properly by working with runes in Go or simply using Python’s string slicing.