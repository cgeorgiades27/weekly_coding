def intersection(list1, list2, list3):
  # Fill this in.
  possibleIntersections = list1
  for index in range(len(list1)):
  
    currentGuess= possibleIntersections[index]
    for j in range(len(list2)):
      if currentGuess> list2[j]:
        continue
      
      if(currentGuess== list2[j]):
        for k in range(len(list3)):
          if(currentGuess> list3[k]):
            continue
      
          if currentGuess==list3[k]:
            return currentGuess
      
      

      
  

  pass
  


def main():
    print("Hello World!")
    print(intersection([1, 2, 3, 6], [2, 4, 6, 8], [3, 5, 6]))
    # [6]

if __name__ == "__main__":
    main()
