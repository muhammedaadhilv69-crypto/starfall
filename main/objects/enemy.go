components {
  id: "enemy"
  component: "/main/scripts/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemyBlack2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/atlases/game.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemies\"\n"
  "mask: \"player\"\n"
  "mask: \"bullets\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"enemycolobj\"\n"
  "  }\n"
  "  data: 48.419834\n"
  "  data: 36.22435\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "enemy_bullet_factory"
  type: "factory"
  data: "prototype: \"/main/objects/enemy_laser.go\"\n"
  ""
}
embedded_components {
  id: "health_bg"
  type: "sprite"
  data: "default_animation: \"health_bar_bg\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/atlases/game.atlas\"\n"
  "}\n"
  ""
  position {
    y: 80.0
    z: -0.1
  }
}
embedded_components {
  id: "health_fill_factory"
  type: "factory"
  data: "prototype: \"/main/objects/health_fill.go\"\n"
  ""
}
