# Naive
class listsum:
	"""Creates a list sum obj and sets properties,
  namely 'sum' which is the sum of both input lists
  returned as a list"""
	
	def __init__(self, list1, list2):
		self.addendA = int(''.join([str(i) for i in list1[::-1]]))
		self.addendB = int(''.join([str(i) for i in list2[::-1]]))
		self.sum = [i for i in str(self.addendA+self.addendB)[::-1]]


# Test
tl1 = [2, 8, 5, 1]
tl2 = [2, 1]		
listint = listsum(tl1,tl2)

print(listint.sum)
