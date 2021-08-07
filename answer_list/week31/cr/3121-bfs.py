from collections import deque


class ComboLock:
    """Class used to represent a three wheel combo lock"""

    def __init__(self, combo, xplist):
        self.combo = str(combo)
        self.setxp = {str(x) for x in xplist}
        self.setxp.add('000')
        self.queue = deque([['000']])
        self.combolist = []

    def check_xp(self, com_check):
        """
        Class method to check if a combo is in the dead end list.
        If it is not it adds it to the queue and the dead end
        so it wont be visited again
        """

        if com_check not in self.setxp:
            self.setxp.add(com_check)
            self.combolist.append(com_check)
            return

    @staticmethod
    def turn_up(wheel, cur_combo):
        """
        Class method to increment a wheel by one.
        Static because it doesn't change the class obj (self)
        """

        lst_combo = [x for x in cur_combo]
        wh_val = int(lst_combo[wheel])
        wh_val += 1
        mod = wh_val % 10  # Mod sets val to 0 if its 10 i.e. 9+1
        lst_combo[wheel] = str(mod)
        new_combo = ''.join(lst_combo)
        return new_combo

    @staticmethod
    def turn_down(wheel, cur_combo):
        """
        Class method to increment a wheel by one.
        Static because it doesn't change the class obj (self)
        """

        lst_combo = [x for x in cur_combo]
        wh_val = int(lst_combo[wheel])
        wh_val -= 1
        mod = wh_val % 10  # sets val to 9 if its -1 i.e. 1-1
        lst_combo[wheel] = str(mod)
        new_combo = ''.join(lst_combo)
        return new_combo

    def turn_wheels(self):
        """
        Class method to count the least amount of rotations
        needed to reach the combo or None if not possible.

        Iterates through a queue and finds all possible combos.
        Each loop builds a list of all combos possible at that level.
        """

        wheels = [2, 1, 0]
        turn_count = 0
        while self.queue:
            cur_queue = self.queue.popleft()

            while cur_queue:
                cur_combo = cur_queue.pop()

                if cur_combo == self.combo:
                    return turn_count

                else:
                    for wheel in wheels:
                        up = self.turn_up(wheel, cur_combo)
                        self.check_xp(up)
                        down = self.turn_down(wheel, cur_combo)
                        self.check_xp(down)

            self.queue.append(self.combolist.copy())
            self.combolist.clear()

            turn_count += 1
        return None


# Test
shallow = ComboLock(222, [999])
deep = ComboLock(555, [999, 111, 100, 900, 816])

print(shallow.turn_wheels())
print(deep.turn_wheels())
