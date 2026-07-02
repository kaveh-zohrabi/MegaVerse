package dev.megaverse.social.controller;

import dev.megaverse.social.model.Comment;
import dev.megaverse.social.model.Post;
import dev.megaverse.social.service.SocialService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping
public class SocialController {

    private final SocialService socialService;

    public SocialController(SocialService socialService) {
        this.socialService = socialService;
    }

    @GetMapping("/health")
    public ResponseEntity<Map<String, String>> health() {
        return ResponseEntity.ok(Map.of("status", "healthy", "service", "social-service"));
    }

    @PostMapping("/posts")
    public ResponseEntity<Post> createPost(@RequestHeader("X-User-ID") String userId,
                                           @RequestBody Map<String, String> body) {
        Post post = socialService.createPost(userId, body.get("content"), body.get("media_url"));
        return ResponseEntity.status(HttpStatus.CREATED).body(post);
    }

    @GetMapping("/posts/{id}")
    public ResponseEntity<Post> getPost(@PathVariable String id) {
        return ResponseEntity.ok(socialService.getPost(id));
    }

    @GetMapping("/users/{userId}/posts")
    public ResponseEntity<List<Post>> getUserPosts(@PathVariable String userId) {
        return ResponseEntity.ok(socialService.getUserPosts(userId));
    }

    @GetMapping("/feed")
    public ResponseEntity<List<Post>> getFeed() {
        return ResponseEntity.ok(socialService.getFeed());
    }

    @DeleteMapping("/posts/{id}")
    public ResponseEntity<Void> deletePost(@PathVariable String id,
                                           @RequestHeader("X-User-ID") String userId) {
        socialService.deletePost(id, userId);
        return ResponseEntity.noContent().build();
    }

    @PostMapping("/posts/{postId}/comments")
    public ResponseEntity<Comment> addComment(@PathVariable String postId,
                                              @RequestHeader("X-User-ID") String userId,
                                              @RequestBody Map<String, String> body) {
        Comment comment = socialService.addComment(postId, userId, body.get("content"));
        return ResponseEntity.status(HttpStatus.CREATED).body(comment);
    }

    @GetMapping("/posts/{postId}/comments")
    public ResponseEntity<List<Comment>> getComments(@PathVariable String postId) {
        return ResponseEntity.ok(socialService.getComments(postId));
    }
}
