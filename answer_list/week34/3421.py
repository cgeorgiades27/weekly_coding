from collections import deque


class Playlist:

    def __init__(self, playlist_graph):
        self.graph = playlist_graph
        self.songs = list(self.graph.keys())
        self.playlist = []

    def make_playlist(self):
        stack = [self.songs.pop()]
        seen = set()
        while stack:
            song = stack.pop()

            if song in seen:
                song = stack.pop()

            elif len(self.graph[song]) == 0:
                self.playlist.append(song)
                seen.add(song)

            else:
                stack.extend(self.graph[song])
                seen.add(song)


input_graph = {1: [2], 7: [1, 2, 6],
               3: [1, 7], 2: [], 6: [2, 1],
               9: [2, 1, 6, 7, 3], 5: [3, 9]}
mega_playlist = Playlist(input_graph)
print(mega_playlist.make_playlist())
