# Trip Kata
Given a list of source and destinations make a function that returns true if there is a route from a source and destination and the number 
connections to get from the source to destination.

For example
['san franscico', 'san diego']
['san diego', 'las vegas']
['san diego', 'los angeles']
['las vegas', 'salt lake city']
['las vegas', 'phoenix']
['phoenix', 'san antonio']
['salt lake city', 'denver']
['denver', 'indianapolis']
['las vegas', 'boise']
['boise', 'alberta']

Where the source = 'san franscico' and the destination = 'denver' should return
['san franscico', 'san diego']
['san diego', 'las vegas']
['las vegas', 'salt lake city']
['salt lake city', 'denver']
with 4 connections
